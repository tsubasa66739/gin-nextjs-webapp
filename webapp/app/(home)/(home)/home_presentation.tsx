"use client";

import { NoteResponse } from "@/schema/note_response";
import { Card } from "@radix-ui/themes";

interface HomePresentationProps {
  notes: NoteResponse[];
}

export default function HomePresentation({ notes }: HomePresentationProps) {
  return (
    <div className="px-3">
      <Card className="px-3 w-full">
        <table className="table-auto border-collapse w-full text-left">
          <thead className="border-b">
            <tr>
              <th>タイトル</th>
              <th>本文</th>
              <th>更新日時</th>
            </tr>
          </thead>
          <tbody>
            {notes.map((n: NoteResponse) => {
              return (
                <tr key={n.id}>
                  <td>{n.title}</td>
                  <td>{n.body}</td>
                  <td>{n.updatedAt}</td>
                </tr>
              );
            })}
          </tbody>
        </table>
      </Card>
    </div>
  );
}
