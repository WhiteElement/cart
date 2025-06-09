import { Component, computed, inject, Signal, WritableSignal, signal } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { ShoppinglistService } from '../shoppinglist.service';
import { ShoppingItemService } from '../shopping-item.service';
import { ShoppingList } from '../models/shopping-list';
import { ShoppingItem } from '../models/shopping-item';
import { DatePipe } from '@angular/common';
import { FormsModule } from '@angular/forms';
import { toSignal } from '@angular/core/rxjs-interop';

@Component({
  selector: 'app-detail',
  standalone: true,
  imports: [DatePipe, FormsModule],
  templateUrl: './detail.component.html',
  styleUrl: './detail.component.css'
})
export class DetailComponent {
  private route = inject(ActivatedRoute);
  private shoppinglistService = inject(ShoppinglistService);
  private shoppingItemService = inject(ShoppingItemService);

  private paramMap = toSignal(this.route.paramMap);
  private id: Signal<string | null | undefined> = computed(() => this.paramMap()?.get('id'));

  //private listFromService$ = this.shoppinglistService.getOne(this.id()!);
  list = this.shoppinglistService.activeList;
  //this.shoppinglistService

  //list = computed(() => {
  //  this.refreshTrigger();
  //  return this.listFromService$;
  //});
  //
  activeItems = computed(() => { return this.list()?.Items?.filter(i => !i.Checked); });
  checkedItems = computed(() => { return this.list()?.Items?.filter(i => i.Checked); });

  toggleNew = signal(false);
  newItemInput = signal('');

  constructor() {
    const id_ = this.id();
    if (id_) {
      this.shoppinglistService.updateList(id_);
    }
  }

  showNew(): void {
    if (this.toggleNew()) {
      this.toggleNew.set(false);
    } else {
      this.toggleNew.set(true);
    }
  }

  createNewItem(): void {
    console.log('in method');
    if (this.newItemInput() !== '') {
      const id = this.list()?.Id;
      console.log('id', id);
      if (id) {
        this.shoppingItemService.newItem(this?.newItemInput()!, id, false).subscribe(res => {
          if (res) {
            this.newItemInput.set('');
            this.toggleNew.set(false);
            this.shoppinglistService.updateList(id.toString());
          }
        });
      }
    }

  }

  toggleCheck(item: ShoppingItem): void {
    if (item.Checked) {
      item.Checked = false;
    } else {
      item.Checked = true;
    }

    //this.list.update(l => {
    //  l?.Items?.push(item);
    //  return l;
    //});

    this.shoppingItemService.patchItem(item).subscribe(res => {
      const statusCode = res.status.toString();
      if (statusCode.startsWith("2")) {
        //this.list.update(l => {
        //  if (l) {
        //    l.Updated = new Date();
        //  }
        //  return l;
        //});
      } else {
        console.error("Error patching Item", res.body);
      }
    });
  }

  delete(itemId: number | null) {
    if (itemId) {
      //this.shoppingItemService.deleteItem(itemId).subscribe(res => {
      //  this.checkedItems = this.checkedItems.filter(i => i.Id !== itemId);
      //  this.activeItems = this.activeItems.filter(i => i.Id !== itemId);
      //  this.list.Updated = new Date();
      //  if (this.list.Items) {
      //    this.list.Items = this.list.Items.filter(i => i.Id !== itemId);
      //  }
      //});
    }
  }
}
