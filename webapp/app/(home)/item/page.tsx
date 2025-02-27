import { ItemResponse } from "@/schema/item_response";
import ItemPresentation from "./item_presentation";

const url = `${process.env.API_URL}/api/item`;

export default async function Item() {
  const res = await fetch(url);
  const item = (await res.json()) as ItemResponse[];

  return <ItemPresentation items={item} />;
}
