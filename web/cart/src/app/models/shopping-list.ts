import { ShoppingItem } from "./shopping-item";

export interface ShoppingList {
  Id: number | null,
  Name: string,
  Items: ShoppingItem[] | null,
  Archived: boolean | null,
  Created: Date | null,
  Updated: Date | null
}
