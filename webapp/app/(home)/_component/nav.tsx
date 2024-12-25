"use client";

import { Flex, IconButton, Tooltip } from "@radix-ui/themes";
import { usePathname, useRouter } from "next/navigation";
import { IconType } from "react-icons";
import {
  AiOutlineCalendar,
  AiOutlineHome,
  AiOutlineLineChart,
  AiOutlineSetting,
} from "react-icons/ai";

type NavItemProps = {
  path: string;
  name: string;
  icon: IconType;
};

function NavItem({ path, name, icon }: NavItemProps) {
  const router = useRouter();
  const pathname = usePathname();

  return (
    <Tooltip content={name} delayDuration={1000} side="right">
      <IconButton
        variant="ghost"
        className={
          pathname === path
            ? "text-neutral-950 cursor-pointer"
            : "text-neutral-400 hover:text-neutral-950 cursor-pointer"
        }
        onClick={() => router.push(path)}
      >
        {icon({ size: 24 })}
      </IconButton>
    </Tooltip>
  );
}

export default function Nav() {
  return (
    <Flex
      direction="column"
      gap="6"
      align="center"
      className="w-10 py-2"
      height="100%"
    >
      <Flex direction="column" gap="5" height="100%">
        <NavItem path="/" name="ホーム" icon={AiOutlineHome} />
        <NavItem path="/calendar" name="カレンダー" icon={AiOutlineCalendar} />
        <NavItem path="/analytics" name="データ" icon={AiOutlineLineChart} />
        <div className="flex-grow" />
        <NavItem path="/setting" name="設定" icon={AiOutlineSetting} />
      </Flex>
    </Flex>
  );
}
