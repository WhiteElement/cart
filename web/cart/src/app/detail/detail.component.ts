import { Component, OnInit } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { ShoppinglistService } from '../shoppinglist.service';
import { ShoppingItemService } from '../shopping-item.service';
import { ShoppingList } from '../models/shopping-list';
import { ShoppingItem } from '../models/shopping-item';
import { DatePipe } from '@angular/common';
import { FormsModule } from '@angular/forms';

@Component({
  selector: 'app-detail',
  standalone: true,
  imports: [DatePipe, FormsModule],
  templateUrl: './detail.component.html',
  styleUrl: './detail.component.css'
})
export class DetailComponent implements OnInit {
  list: ShoppingList = {} as ShoppingList;
  activeItems: ShoppingItem[] = [];
  checkedItems: ShoppingItem[] = [];

  toggleNew: boolean = false;
  newItemInput: string = '';

  constructor(private route: ActivatedRoute, private shoppinglistService: ShoppinglistService, private shoppingItemService: ShoppingItemService) { }

  ngOnInit(): void {
    this.route.paramMap.subscribe(params => {
      const id = params.get('id');
      if (id) {
        this.shoppinglistService.getOne(id).subscribe(res => {
          const statusCode = res.status.toString();
          if (statusCode.startsWith("2")) {
            if (res.body) {
              this.list = res.body;

              if (this.list.Items) {
                this.activeItems = this.list.Items.filter(i => !i.Checked);
                this.checkedItems = this.list.Items.filter(i => i.Checked);
              }
            }
          }
        })
      }
    });
  }

  showNew(): void {
    if (this.toggleNew) {
      this.toggleNew = false;
    } else {
      this.toggleNew = true;
    }
  }

  createNewItem(): void {
    if (this.newItemInput !== '') {
      if (this.list.Id) {
        this.shoppingItemService.newItem(this.newItemInput, this.list.Id, false).subscribe(res => {
          const statusCode = res.status.toString();
          if (statusCode.startsWith("2")) {
            this.newItemInput = '';
            this.toggleNew = false;
            this.ngOnInit();
          }
        });
      }
    }

  }

  toggleCheck(item: ShoppingItem): void {
    if (item.Checked) {
      item.Checked = false;
      this.activeItems.push(item);
      this.checkedItems = this.checkedItems.filter(i => i.Id !== item.Id);
    } else {
      item.Checked = true;
      this.checkedItems.push(item);
      this.activeItems = this.activeItems.filter(i => i.Id !== item.Id);
    }

    this.shoppingItemService.patchItem(item).subscribe(res => {
      const statusCode = res.status.toString();
      if (statusCode.startsWith("2")) {
        this.list.Updated = new Date();
      } else {
        console.error("Error patching Item", res.body);
      }
    });
  }

  delete(itemId: number | null) {
    if (itemId) {
      this.shoppingItemService.deleteItem(itemId).subscribe(res => {
        this.checkedItems = this.checkedItems.filter(i => i.Id !== itemId);
        this.activeItems = this.activeItems.filter(i => i.Id !== itemId);
        this.list.Updated = new Date();
        if (this.list.Items) {
          this.list.Items = this.list.Items.filter(i => i.Id !== itemId);
        }
      });
    }
  }
}
