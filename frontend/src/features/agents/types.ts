export type AgentType = 'General' | 'Custom'
export type AgentStatus = 'Active' | 'Inactive' | 'Inbound' | 'Outbound'
export type AgentVoice = {
  name: string
  language: string
  settings?: Record<string, unknown>
}

export interface Agent {
  id: string
  name: string
  phoneNumber: string
  greeting: string
  instructions: string
  type: AgentType
  status: AgentStatus
  avatar?: string
  voice?: AgentVoice
  metrics?: {
    callsMade: number
    avgCostPerCall: number
    avgTalkTime: number
    totalTalkTime: number
    satisfactionScore: string
    firstCallResolution: string
  }
  recordings?: Array<{
    id: string
    customerNumber: string
    timestamp: string
    duration: string
    url: string
  }>
}
