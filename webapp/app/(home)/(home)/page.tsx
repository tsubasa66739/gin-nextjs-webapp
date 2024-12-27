import { NoteResponse } from "@/schema/note_response";
import HomePresentation from "./home_presentation";

export default async function Home() {
  const res = await fetch("http://localhost:8080/api/note");
  const notes = (await res.json()) as NoteResponse[];

  return <HomePresentation notes={notes} />;
}
