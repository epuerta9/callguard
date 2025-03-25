"use client";

import { Bell, Search, User } from "lucide-react";

interface HeaderProps {
  title: string;
}

export function Header({ title }: HeaderProps) {
  return (
    <header className="bg-white dark:bg-slate-800 border-b border-slate-200 dark:border-slate-700 h-16 flex items-center justify-between px-6">
      <h1 className="text-xl font-semibold">{title}</h1>
      
      <div className="flex items-center gap-4">
        <div className="relative">
          <Search className="absolute left-3 top-1/2 transform -translate-y-1/2 text-slate-400" size={18} />
          <input 
            type="text" 
            placeholder="Search..." 
            className="pl-10 pr-4 py-2 rounded-md border border-slate-200 dark:border-slate-600 bg-slate-50 dark:bg-slate-700 w-64"
          />
        </div>
        
        <button className="relative">
          <Bell size={20} />
          <span className="absolute -top-1 -right-1 bg-red-500 text-white text-xs rounded-full w-4 h-4 flex items-center justify-center">
            3
          </span>
        </button>
        
        <div className="flex items-center gap-2">
          <div className="w-8 h-8 rounded-full bg-slate-200 flex items-center justify-center">
            <User size={16} />
          </div>
          <span className="text-sm font-medium">Admin</span>
        </div>
      </div>
    </header>
  );
} 