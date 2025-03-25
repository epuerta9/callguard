import {
  Sidebar,
  SidebarContent,
  SidebarFooter,
  SidebarHeader,
  SidebarRail,
} from "@/components/ui/sidebar";
import { NavGroup } from "@/components/layout/nav-group";
import { NavUser } from "@/components/layout/nav-user";
import { sidebarData } from "./data/sidebar-data";
import { useAuth } from "@/hooks/use-auth";
import OctavioLogo from "@/assets/octavio_icon_black_white.svg?react";

export function AppSidebar({ ...props }: React.ComponentProps<typeof Sidebar>) {
  const { user } = useAuth();
  return (
    <Sidebar collapsible="icon" variant="floating" {...props}>
      <SidebarHeader>
        <div className="group flex items-center gap-2 p-2">
          <span className="flex-1 text-lg font-semibold leading-tight group-data-[collapsible=icon]:hidden">
            CallGuard
          </span>
        </div>
      </SidebarHeader>
      <SidebarContent>
        {sidebarData.navGroups.map((props) => (
          <NavGroup key={props.title} {...props} />
        ))}
      </SidebarContent>
      <SidebarFooter>
        <NavUser user={user || {}} />
      </SidebarFooter>
      <SidebarRail />
    </Sidebar>
  );
}
