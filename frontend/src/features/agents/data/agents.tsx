import { Agent } from "../types";

export const agents: Agent[] = [
  {
    id: "1",
    name: "Mando's grandma",
    phoneNumber: "(713)-587-6237",
    avatar: "/avatars/agent1.jpg",
    type: "General",
    status: "Inbound",
    greeting: "Hello from Barber Studio",
    instructions:
      'You are an AI assistant acting as the virtual receptionist for a busy barbershop.\n\nYou\'re answering the phone for "{{CompanyName}}", a barbershop that offers haircuts, shaves, and other grooming services.\n\nYour tasks include providing information about services and products offered, such as their costs and the time required for each service. Additionally, you can schedule appointments, reschedule existing ones, or cancel appointments upon request.',
    voice: {
      name: "Chris",
      language: "English (US)",
    },
    metrics: {
      callsMade: 145,
      avgCostPerCall: 0.12,
      avgTalkTime: 3.5,
      totalTalkTime: 507.5,
      satisfactionScore: "4.8/5",
      firstCallResolution: "92%",
    },
    recordings: [
      {
        id: "rec1",
        customerNumber: "(415) 555-0123",
        timestamp: "2024-03-25T08:30:00Z",
        duration: "3:45",
        url: "#",
      },
      {
        id: "rec2",
        customerNumber: "(415) 555-0124",
        timestamp: "2024-03-25T09:15:00Z",
        duration: "2:30",
        url: "#",
      },
    ],
  },
];
