import { NoteResponse } from "@/schema/note_response";
import HomePresentation from "./home_presentation";

const url = "http://localhost:8080";

export default async function Home() {
  const res = await fetch(`${url}/api/note`);
  const notes = (await res.json()) as NoteResponse[];

  return <HomePresentation notes={notes} />;
}
