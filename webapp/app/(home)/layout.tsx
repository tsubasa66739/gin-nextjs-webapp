import { ReactNode } from "react";
import Nav from "./_component/nav";

export default function DashboardLayout({
  children,
}: Readonly<{
  children: ReactNode;
}>) {
  return (
    <div className="flex w-screen h-screen bg-slate-50">
      <nav className="p-2 bg-white border-r">
        <Nav />
      </nav>
      <main className="py-4 w-full">{children}</main>
    </div>
  );
}
