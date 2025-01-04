import { NoteResponse } from "@/schema/note_response";
import HomePresentation from "./home_presentation";

const url = `${process.env.API_URL}/api/note`;

export default async function Home() {
  const res = await fetch(url);
  const notes = (await res.json()) as NoteResponse[];

  return <HomePresentation notes={notes} />;
}
