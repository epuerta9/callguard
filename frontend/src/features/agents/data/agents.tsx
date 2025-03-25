export interface Agent {
  id: string;
  name: string;
  phoneNumber: string;
  avatar?: string;
  type: "General" | "Custom";
  status: "Inbound" | "Outbound";
}

export const agents: Agent[] = [
  {
    id: "1",
    name: "Mando's Grandma",
    phoneNumber: "(713)-587-6237",
    avatar: "/avatars/agent1.jpg",
    type: "General",
    status: "Inbound",
  },
];
