"use client";

import { Button, Tooltip } from "@nextui-org/react";
import { usePathname, useRouter } from "next/navigation";
import { IconType } from "react-icons";
import {
  AiOutlineCalendar,
  AiOutlineHome,
  AiOutlineLineChart,
  AiOutlineSetting,
} from "react-icons/ai";

interface NavItemProps {
  path: string;
  name: string;
  icon: IconType;
}

function NavItem({ path, name, icon }: NavItemProps) {
  const router = useRouter();
  const pathname = usePathname();

  return (
    <Tooltip
      aria-label={name}
      content={name}
      placement="right"
      size="sm"
      delay={700}
      closeDelay={100}
    >
      <Button
        aria-label={name}
        isIconOnly
        color="primary"
        variant="light"
        className={
          pathname === path
            ? "text-neutral-950 bg-gray-200"
            : "text-neutral-400 hover:text-neutral-950"
        }
        onPress={() => router.push(path)}
      >
        {icon({ size: 24 })}
      </Button>
    </Tooltip>
  );
}

export default function Nav() {
  return (
    <div className="flex flex-col gap-5 align-middle w-10 py-2 h-full">
      <NavItem path="/" name="ホーム" icon={AiOutlineHome} />
      <NavItem path="/calendar" name="カレンダー" icon={AiOutlineCalendar} />
      <NavItem path="/analytics" name="データ" icon={AiOutlineLineChart} />
      <div className="flex-grow" />
      <NavItem path="/setting" name="設定" icon={AiOutlineSetting} />
    </div>
  );
}
