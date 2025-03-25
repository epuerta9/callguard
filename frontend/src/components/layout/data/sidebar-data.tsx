import {
  IconDashboard,
  IconUsers,
  IconPhone,
  IconHistory,
  IconSettings,
  IconHelp,
  IconShieldLock,
  IconUserCog,
  IconBell,
} from "@tabler/icons-react";
import { type SidebarData } from "../types";

export const sidebarData: SidebarData = {
  user: {
    name: "Demo User",
    email: "demo@callguard.ai",
    avatar: "/avatars/default.jpg",
  },
  navGroups: [
    {
      title: "Main",
      items: [
        {
          title: "Dashboard",
          url: "/",
          icon: IconDashboard,
        },
        {
          title: "Voice Agents",
          url: "/agents",
          icon: IconUsers,
        },
        {
          title: "Phone Numbers",
          url: "/phone-numbers",
          icon: IconPhone,
        },
        {
          title: "Call History",
          url: "/call-history",
          icon: IconHistory,
        },
      ],
    },
    {
      title: "Settings",
      items: [
        {
          title: "Account",
          url: "/settings/account",
          icon: IconUserCog,
        },
        {
          title: "Security",
          url: "/settings/security",
          icon: IconShieldLock,
        },
        {
          title: "Notifications",
          url: "/settings/notifications",
          icon: IconBell,
        },
      ],
    },
    {
      title: "Support",
      items: [
        {
          title: "Help Center",
          url: "/help",
          icon: IconHelp,
        },
      ],
    },
  ],
};
