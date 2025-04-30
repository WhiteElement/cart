import { ShoppingItem } from "./shopping-item";

export interface ShoppingList {
  Id: number,
  Name: string,
  Items: ShoppingItem[],
  Archived: boolean,
  Created: Date,
  Updated: Date
}
