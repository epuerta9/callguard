import { Card } from "@/components/ui/card";
import { Agent } from "../types";
import { CalendarIcon, PhoneIcon } from "lucide-react";
import { Button } from "@/components/ui/button";
import { format, parseISO } from "date-fns";

interface AgentActivityProps {
  agent: Agent;
}

interface MetricCardProps {
  title: string;
  value: string | number;
  info?: string;
}

function MetricCard({ title, value, info }: MetricCardProps) {
  return (
    <div className="rounded-lg border p-4">
      <div className="text-sm text-muted-foreground">{title}</div>
      <div className="mt-1 text-2xl font-bold">{value}</div>
      {info && <div className="text-xs text-muted-foreground mt-1">{info}</div>}
    </div>
  );
}

export function AgentActivity({ agent }: AgentActivityProps) {
  const today = new Date();

  return (
    <div className="space-y-6">
      <div className="flex items-center justify-between">
        <div className="flex items-center gap-2">
          <h3 className="text-lg font-medium">Agent Activity</h3>
          <Button variant="outline" size="sm" className="gap-2">
            <CalendarIcon className="h-4 w-4" />
            {format(today, "EEEE, dd MMM yyyy")}
          </Button>
        </div>
      </div>

      <div className="grid gap-4 sm:grid-cols-2 lg:grid-cols-3">
        <MetricCard title="Calls made" value={agent.metrics?.callsMade || 0} />
        <MetricCard
          title="Avg. cost per call"
          value={`$${agent.metrics?.avgCostPerCall?.toFixed(2) || "0.00"}`}
        />
        <MetricCard
          title="Avg. talk time"
          value={`${agent.metrics?.avgTalkTime || "0.0"} min`}
        />
        <MetricCard
          title="Total talk time"
          value={`${agent.metrics?.totalTalkTime || "0.0"} min`}
        />
        <MetricCard
          title="Satisfaction score"
          value={agent.metrics?.satisfactionScore || "N/A"}
        />
        <MetricCard
          title="First call resolution"
          value={agent.metrics?.firstCallResolution || "N/A"}
        />
      </div>

      <div className="space-y-4">
        <h4 className="text-lg font-medium">Recent Recordings</h4>
        <div className="space-y-4">
          {agent.recordings && agent.recordings.length > 0 ? (
            agent.recordings.map((recording: { id: string; customerNumber: string; timestamp: string; duration: string }) => (
              <Card key={recording.id} className="p-4">
                <div className="flex items-center justify-between">
                  <div className="flex items-center gap-4">
                    <div className="rounded-full bg-primary/10 p-2">
                      <PhoneIcon className="h-4 w-4 text-primary" />
                    </div>
                    <div>
                      <div className="font-medium">{recording.customerNumber}</div>
                      <div className="text-sm text-muted-foreground">
                        {format(parseISO(recording.timestamp), "MMM d, yyyy h:mm a")}
                      </div>
                    </div>
                  </div>
                  <div className="flex items-center gap-4">
                    <div className="text-sm text-muted-foreground">
                      {recording.duration}
                    </div>
                    <Button variant="ghost" size="sm">
                      Play
                    </Button>
                  </div>
                </div>
              </Card>
            ))
          ) : (
            <div className="text-sm text-muted-foreground">
              No recordings available.
            </div>
          )}
        </div>
      </div>
    </div>
  );
}
