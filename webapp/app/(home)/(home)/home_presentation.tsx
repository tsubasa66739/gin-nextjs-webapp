"use client";

import { NoteResponse } from "@/schema/note_response";
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

interface HomePresentationProps {
  notes: NoteResponse[];
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

export default function HomePresentation({ notes }: HomePresentationProps) {
  return (
    <div className="px-3">
      <div className="flex justify-end mb-4 gap-2">
        <Button color="default" size="sm">
          Other control
        </Button>
        <Button color="primary" size="sm">
          New Note
        </Button>
      </div>
      <Table aria-label="Note data table">
        <TableHeader>
          <TableColumn>タイトル</TableColumn>
          <TableColumn>本文</TableColumn>
          <TableColumn>更新日時</TableColumn>
          <TableColumn>{""}</TableColumn>
        </TableHeader>
        <TableBody>
          {notes.map((n: NoteResponse) => {
            return (
              <TableRow key={n.id} className="border-b">
                <TableCell>{n.title}</TableCell>
                <TableCell>{n.body}</TableCell>
                <TableCell>{formatDatetime(n.updatedAt)}</TableCell>
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
