"use client";

import Link from "next/link";
import { Home, Users, Phone, AlertTriangle } from "lucide-react";

interface SidebarProps {
  className?: string;
}

export function Sidebar({ className }: SidebarProps) {
  return (
    <div className={`${className} bg-slate-900 text-white h-screen w-64 p-4 fixed left-0 top-0`}>
      <div className="mb-8">
        <h1 className="text-2xl font-bold">CallGuard</h1>
        <p className="text-slate-400 text-sm">Voice agent protection</p>
      </div>

      <nav className="space-y-2">
        <Link 
          href="/dashboard"
          className="flex items-center gap-3 px-3 py-2 rounded-md hover:bg-slate-800 transition-colors"
        >
          <Home size={18} />
          <span>Dashboard</span>
        </Link>

        <Link 
          href="/dashboard/agents"
          className="flex items-center gap-3 px-3 py-2 rounded-md hover:bg-slate-800 transition-colors"
        >
          <Users size={18} />
          <span>Agents</span>
        </Link>

        <Link 
          href="/dashboard/calls"
          className="flex items-center gap-3 px-3 py-2 rounded-md hover:bg-slate-800 transition-colors"
        >
          <Phone size={18} />
          <span>Call History</span>
        </Link>

        <Link 
          href="/dashboard/alerts"
          className="flex items-center gap-3 px-3 py-2 rounded-md hover:bg-slate-800 transition-colors"
        >
          <AlertTriangle size={18} />
          <span>Alerts</span>
        </Link>
      </nav>
    </div>
  );
} 