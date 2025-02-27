"use client";

import { ItemResponse } from "@/schema/item_response";
import { Button } from "@nextui-org/button";
import {
  Table,
  TableBody,
  TableCell,
  TableColumn,
  TableHeader,
  TableRow,
} from "@nextui-org/table";
import { AiOutlineMore } from "react-icons/ai";

interface ItemPresentationProps {
  items: ItemResponse[];
}

function formatDatetime(dateStr: string) {
  const datetime = new Date(dateStr);
  const year = datetime.getFullYear();
  const month = datetime.getMonth() + 1;
  const date = datetime.getDate();
  const hour = String(datetime.getHours()).padStart(2, "0");
  const minute = String(datetime.getMinutes()).padStart(2, "0");
  return `${year}/${month}/${date} ${hour}:${minute}`;
}

export default function ItemPresentation({ items }: ItemPresentationProps) {
  return (
    <div className="px-3">
      <div className="flex justify-end mb-4 gap-2">
        <Button color="primary" size="sm" variant="bordered">
          Other control
        </Button>
        <Button color="primary" size="sm">
          New item
        </Button>
      </div>
      <Table aria-label="item data table">
        <TableHeader>
          <TableColumn>名前</TableColumn>
          <TableColumn>値段</TableColumn>
          <TableColumn>更新日時</TableColumn>
          <TableColumn>{""}</TableColumn>
        </TableHeader>

        <TableBody emptyContent={"No data."}>
          
          {items.map((i: ItemResponse) => {
            return (
              <TableRow key={i.id} className="border-b hover:bg-gray-100">
                <TableCell>{i.name}</TableCell>
                <TableCell>{i.price}</TableCell>
                <TableCell>{formatDatetime(i.updatedAt)}</TableCell>
                <TableCell className="text-right">
                  <Button isIconOnly variant="light" size="sm">
                    <AiOutlineMore className="text-xl" />
                  </Button>
                </TableCell>
              </TableRow>
            );
          })}
        </TableBody>
      </Table>
    </div>
  
  );
}
