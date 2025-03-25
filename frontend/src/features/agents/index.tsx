import { useState, useEffect } from "react";
import { useNavigate } from "@tanstack/react-router";
import {
  IconPlus,
  IconSearch,
  IconUser,
} from "@tabler/icons-react";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogFooter,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from "@/components/ui/dialog";
import { Label } from "@/components/ui/label";
import { Header } from "@/components/layout/header";
import { Main } from "@/components/layout/main";
import { ProfileDropdown } from "@/components/profile-dropdown";
import { ThemeSwitch } from "@/components/theme-switch";
import { Avatar, AvatarFallback } from "@/components/ui/avatar";
import { Badge } from "@/components/ui/badge";
import { Card, CardHeader } from "@/components/ui/card";
import { userMetadataApi, type Agent } from "@/lib/api/user-metadata";
import { toast } from "sonner";

export default function Agents() {
  const navigate = useNavigate();
  const [sort] = useState("ascending");
  const [agents, setAgents] = useState<Agent[]>([]);
  const [isLoading, setIsLoading] = useState(true);

  useEffect(() => {
    loadAgents();
  }, []);

  const loadAgents = async () => {
    try {
      const metadata = await userMetadataApi.get();
      const agentsList = Object.entries(metadata.agents || {}).map(([id, agent]) => ({
        id,
        ...(agent as Omit<Agent, 'id'>)
      }));
      setAgents(agentsList);
    } catch (_error) {
      toast.error("Failed to load agents");
    } finally {
      setIsLoading(false);
    }
  };

  const handleAgentClick = (id: string) => {
    navigate({ to: "/agents/[id]" as const, params: { id } });
  };

  const [searchTerm, setSearchTerm] = useState("");
  const [isCreateOpen, setIsCreateOpen] = useState(false);
  const [newAgent, setNewAgent] = useState<Omit<Agent, 'id'>>({
    name: "",
    description: "",
    model: "gpt-4",
    systemPrompt: "",
    temperature: 0.7,
    maxTokens: 1000,
    topP: 1,
    frequencyPenalty: 0,
    presencePenalty: 0,
    stopSequences: []
  });

  const filteredAgents = [...agents]
    .sort((a: Agent, b: Agent) =>
      sort === "ascending"
        ? a.name.localeCompare(b.name)
        : b.name.localeCompare(a.name)
    )
    .filter(
      (agent: Agent) =>
        agent.name.toLowerCase().includes(searchTerm.toLowerCase())
    );

  const handleCreateAgent = async () => {
    try {
      await userMetadataApi.update({
        agents: {
          [crypto.randomUUID()]: newAgent
        }
      });
      await loadAgents();
      setIsCreateOpen(false);
      setNewAgent({
        name: "",
        description: "",
        model: "gpt-4",
        systemPrompt: "",
        temperature: 0.7,
        maxTokens: 1000,
        topP: 1,
        frequencyPenalty: 0,
        presencePenalty: 0,
        stopSequences: []
      });
      toast.success("Agent created successfully");
    } catch (_error) {
      toast.error("Failed to create agent");
    }
  };

  return (
    <>
      <Header>
        <div className="flex items-center gap-4">
          <h1 className="text-2xl font-bold">Agents</h1>
          <Badge variant="secondary">Dashboard / Agents</Badge>
        </div>
        <div className="ml-auto flex items-center gap-4">
          <ThemeSwitch />
          <ProfileDropdown />
        </div>
      </Header>

      <Main>
        <div className="flex items-center justify-between mb-6">
          <div>
            <h2 className="text-xl font-semibold">Your agents</h2>
            <p className="text-muted-foreground">
              Click one of your agent to see its setting or analytics
            </p>
          </div>
          <Dialog open={isCreateOpen} onOpenChange={setIsCreateOpen}>
            <DialogTrigger asChild>
              <Button>
                <IconPlus className="mr-2" size={16} />
                Create New
              </Button>
            </DialogTrigger>
            <DialogContent>
              <DialogHeader>
                <DialogTitle>Create New Agent</DialogTitle>
                <DialogDescription>
                  Add a new voice agent to handle your calls
                </DialogDescription>
              </DialogHeader>
              <form
                id="agent-form"
                onSubmit={(e) => {
                  e.preventDefault();
                  handleCreateAgent();
                }}
              >
                <div className="grid gap-4 py-4">
                  <div className="grid gap-2">
                    <Label htmlFor="name">Name</Label>
                    <Input
                      id="name"
                      value={newAgent.name}
                      onChange={(e) =>
                        setNewAgent((prev) => ({
                          ...prev,
                          name: e.target.value,
                        }))
                      }
                      placeholder="Enter agent name"
                      required
                    />
                  </div>
                  <div className="grid gap-2">
                    <Label htmlFor="description">Description</Label>
                    <Input
                      id="description"
                      value={newAgent.description}
                      onChange={(e) =>
                        setNewAgent((prev) => ({
                          ...prev,
                          description: e.target.value,
                        }))
                      }
                      placeholder="Enter agent description"
                      required
                    />
                  </div>
                  <div className="grid gap-2">
                    <Label htmlFor="systemPrompt">System Prompt</Label>
                    <Input
                      id="systemPrompt"
                      value={newAgent.systemPrompt}
                      onChange={(e) =>
                        setNewAgent((prev) => ({
                          ...prev,
                          systemPrompt: e.target.value,
                        }))
                      }
                      placeholder="Enter system prompt"
                      required
                    />
                  </div>
                </div>
                <DialogFooter>
                  <Button
                    type="button"
                    variant="outline"
                    onClick={() => setIsCreateOpen(false)}
                  >
                    Cancel
                  </Button>
                  <Button type="submit">Save changes</Button>
                </DialogFooter>
              </form>
            </DialogContent>
          </Dialog>
        </div>

        <div className="mb-6">
          <div className="relative">
            <IconSearch
              className="absolute left-3 top-1/2 -translate-y-1/2 text-muted-foreground"
              size={18}
            />
            <Input
              placeholder="Search agents..."
              value={searchTerm}
              onChange={(e) => setSearchTerm(e.target.value)}
              className="pl-10 w-[300px]"
            />
          </div>
        </div>

        {isLoading ? (
          <div className="text-center py-8">Loading agents...</div>
        ) : (
          <div className="grid gap-4 md:grid-cols-2 lg:grid-cols-3">
            {filteredAgents.map((agent) => (
              <Card
                key={agent.id}
                className="cursor-pointer hover:shadow-md transition-shadow"
                onClick={() => handleAgentClick(agent.id)}
              >
                <CardHeader className="pb-4">
                  <div className="flex items-center gap-4">
                    <Avatar>
                      <AvatarFallback>
                        <IconUser size={24} />
                      </AvatarFallback>
                    </Avatar>
                    <div>
                      <h3 className="font-medium">{agent.name}</h3>
                      <p className="text-sm text-muted-foreground">
                        {agent.description}
                      </p>
                    </div>
                  </div>
                </CardHeader>
              </Card>
            ))}
          </div>
        )}
      </Main>
    </>
  );
}
