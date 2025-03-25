import Link from "next/link";
import { Plus, Phone, Edit, Trash2 } from "lucide-react";
import { DashboardLayout } from "@/components/layout/dashboard-layout";

export default function AgentsPage() {
  // Mock data for presentation
  const agents = [
    {
      id: "1",
      relativeName: "John Doe",
      phoneNumber: "(555) 123-4567",
      guardrails: ["No financial information", "No personal details", "Verify caller identity"],
      createdAt: new Date().toISOString(),
      updatedAt: new Date().toISOString()
    },
    {
      id: "2",
      relativeName: "Jane Smith",
      phoneNumber: "(555) 987-6543",
      guardrails: ["Block suspicious numbers", "No account changes"],
      createdAt: new Date().toISOString(),
      updatedAt: new Date().toISOString()
    },
    {
      id: "3",
      relativeName: "Robert Johnson",
      phoneNumber: "(555) 555-1212",
      guardrails: ["No personal information", "No financial requests"],
      createdAt: new Date().toISOString(),
      updatedAt: new Date().toISOString()
    }
  ];
  
  return (
    <DashboardLayout title="Agents">
      <div className="mb-6 flex items-center justify-between">
        <h2 className="text-xl font-semibold">Your Agents</h2>
        <Link 
          href="/dashboard/agents/new" 
          className="bg-indigo-600 hover:bg-indigo-700 text-white px-4 py-2 rounded-md inline-flex items-center"
        >
          <Plus size={18} className="mr-1" />
          Create New
        </Link>
      </div>
      
      <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
        {agents.map((agent) => (
          <div key={agent.id} className="bg-white dark:bg-slate-800 rounded-lg shadow p-5 border border-slate-200 dark:border-slate-700">
            <div className="flex items-center justify-between mb-4">
              <div className="flex items-center gap-3">
                <div className="w-12 h-12 rounded-full bg-indigo-100 dark:bg-indigo-900 flex items-center justify-center">
                  <span className="text-indigo-600 dark:text-indigo-300 text-lg font-semibold">
                    {agent.relativeName.charAt(0).toUpperCase()}
                  </span>
                </div>
                <div>
                  <h3 className="font-medium text-lg">{agent.relativeName}</h3>
                  <div className="flex items-center text-slate-500 dark:text-slate-400">
                    <Phone size={16} className="mr-1" />
                    <span>{agent.phoneNumber}</span>
                  </div>
                </div>
              </div>
              
              <div className="flex gap-2">
                <button className="p-2 text-slate-500 dark:text-slate-400 hover:text-indigo-600 dark:hover:text-indigo-400">
                  <Edit size={18} />
                </button>
                <button className="p-2 text-slate-500 dark:text-slate-400 hover:text-red-600 dark:hover:text-red-400">
                  <Trash2 size={18} />
                </button>
              </div>
            </div>
            
            <div className="mt-4">
              <h4 className="text-sm font-medium text-slate-500 dark:text-slate-400 mb-2">Guardrails</h4>
              <div className="flex flex-wrap gap-2">
                {agent.guardrails.map((guardrail, index) => (
                  <span 
                    key={index}
                    className="px-2 py-1 bg-slate-100 dark:bg-slate-700 text-slate-700 dark:text-slate-300 text-xs rounded-full"
                  >
                    {guardrail}
                  </span>
                ))}
              </div>
            </div>
          </div>
        ))}
      </div>
    </DashboardLayout>
  );
} 