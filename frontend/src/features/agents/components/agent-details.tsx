import { Button } from "@/components/ui/button";
import { Card } from "@/components/ui/card";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { Tabs, TabsContent, TabsList, TabsTrigger } from "@/components/ui/tabs";
import { Textarea } from "@/components/ui/textarea";
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "@/components/ui/select";
import { Agent } from "../types";
import { AgentActivity } from "./agent-activity";
import { useState } from "react";

interface AgentDetailsProps {
  agent: Agent;
  onUpdate: (updatedAgent: Agent) => void;
}

export function AgentDetails({ agent, onUpdate }: AgentDetailsProps) {
  const [isTestCallActive, setIsTestCallActive] = useState(false);
  const [testPhoneNumber, setTestPhoneNumber] = useState("");

  const handleTestCall = () => {
    setIsTestCallActive(true);
    // Implement actual call functionality here
    setTimeout(() => {
      setIsTestCallActive(false);
    }, 10000); // 10 seconds timeout for demo
  };

  return (
    <div className="space-y-6">
      {/* Test Agent Card */}
      <Card className="p-6">
        <div className="flex items-center justify-between">
          <div>
            <h3 className="text-lg font-medium">Test your agent</h3>
            <p className="text-sm text-muted-foreground">
              Let your agent call you. Call minutes left: 10
            </p>
          </div>
          <div className="flex items-center gap-2">
            <Input
              value={testPhoneNumber}
              onChange={(e) => setTestPhoneNumber(e.target.value)}
              placeholder="(415) 222-2345"
              className="w-40"
            />
            <Button
              variant="secondary"
              onClick={handleTestCall}
              disabled={isTestCallActive}
            >
              {isTestCallActive ? "Calling..." : "Call me"}
            </Button>
          </div>
        </div>
      </Card>

      {/* Agent Configuration Tabs */}
      <Tabs defaultValue="setup" className="space-y-4">
        <TabsList>
          <TabsTrigger value="setup">Setup</TabsTrigger>
          <TabsTrigger value="knowledge">Knowledge</TabsTrigger>
          <TabsTrigger value="activity">Activity</TabsTrigger>
        </TabsList>

        <TabsContent value="setup" className="space-y-4">
          <div className="grid gap-4">
            <div className="grid gap-2">
              <Label>Voice</Label>
              <Select defaultValue="chris">
                <SelectTrigger>
                  <SelectValue placeholder="Select a voice" />
                </SelectTrigger>
                <SelectContent>
                  <SelectItem value="chris">Chris</SelectItem>
                  <SelectItem value="emma">Emma</SelectItem>
                  <SelectItem value="dave">Dave</SelectItem>
                </SelectContent>
              </Select>
            </div>

            <div className="grid gap-2">
              <Label>Language</Label>
              <Select defaultValue="en-US">
                <SelectTrigger>
                  <SelectValue placeholder="Select a language" />
                </SelectTrigger>
                <SelectContent>
                  <SelectItem value="en-US">English (US)</SelectItem>
                  <SelectItem value="es-ES">Spanish (Spain)</SelectItem>
                </SelectContent>
              </Select>
            </div>

            <div className="grid gap-2">
              <Label>Advanced Voice Settings</Label>
              <Button variant="outline">Configure</Button>
            </div>
          </div>
        </TabsContent>

        <TabsContent value="knowledge" className="space-y-4">
          <div className="grid gap-4">
            <div className="grid gap-2">
              <Label>Greeting</Label>
              <Input
                placeholder="Hello from {Company Name}"
                defaultValue={agent.greeting}
                onChange={(e) =>
                  onUpdate({ ...agent, greeting: e.target.value })
                }
              />
            </div>

            <div className="grid gap-2">
              <Label>Instructions</Label>
              <Textarea
                placeholder="Write instructions for your agent..."
                className="min-h-[150px]"
                defaultValue={agent.instructions}
                onChange={(e) =>
                  onUpdate({ ...agent, instructions: e.target.value })
                }
              />
            </div>
          </div>
        </TabsContent>

        <TabsContent value="activity">
          <AgentActivity agent={agent} />
        </TabsContent>
      </Tabs>
    </div>
  );
}
