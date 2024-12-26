import { ReactNode } from "react";
import Nav from "./_component/nav";
import { Flex } from "@radix-ui/themes";

export default function DashboardLayout({
  children,
}: Readonly<{
  children: ReactNode;
}>) {
  return (
    <Flex className="w-screen h-screen bg-slate-50">
      <nav className="p-2 bg-white border-r">
        <Nav />
      </nav>
      <main className="py-4">{children}</main>
    </Flex>
  );
}
