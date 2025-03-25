import { useParams } from '@tanstack/react-router'
import { AgentDetails } from '@/features/agents/components/agent-details'
import { Main } from '@/components/layout/main'
import { Header } from '@/components/layout/header'
import { ProfileDropdown } from '@/components/profile-dropdown'
import { ThemeSwitch } from '@/components/theme-switch'
import { Button } from '@/components/ui/button'
import { IconArrowLeft } from '@tabler/icons-react'
import { useNavigate } from '@tanstack/react-router'
import { userMetadataApi } from '@/lib/api/user-metadata'
import { useEffect, useState } from 'react'
import { toast } from 'sonner'
import { Agent, AgentType, AgentStatus } from '@/features/agents/types'

export default function AgentDetailsPage() {
  const { id } = useParams({ from: '/_authenticated/agents/[id]' })
  const navigate = useNavigate()
  const [agent, setAgent] = useState<Agent | null>(null)
  const [isLoading, setIsLoading] = useState(true)

  useEffect(() => {
    loadAgent()
  }, [id])

  const loadAgent = async () => {
    try {
      const metadata = await userMetadataApi.get()
      const agents = metadata.agents || {}
      const foundAgent = agents[id]
      if (foundAgent) {
        setAgent({ 
          id,
          name: foundAgent.name || '',
          phoneNumber: foundAgent.phoneNumber || '',
          greeting: foundAgent.greeting || '',
          instructions: foundAgent.instructions || '',
          type: foundAgent.type || 'General' as AgentType,
          status: foundAgent.status || 'Active' as AgentStatus,
          avatar: foundAgent.avatar,
          voice: foundAgent.voice,
          metrics: foundAgent.metrics,
          recordings: foundAgent.recordings
        })
      }
    } catch (_error) {
      toast.error('Failed to load agent')
    } finally {
      setIsLoading(false)
    }
  }

  const handleUpdate = async (updatedAgent: Agent) => {
    try {
      await userMetadataApi.update({
        agents: {
          [id]: {
            name: updatedAgent.name,
            phoneNumber: updatedAgent.phoneNumber,
            greeting: updatedAgent.greeting,
            instructions: updatedAgent.instructions,
            type: updatedAgent.type,
            status: updatedAgent.status,
            avatar: updatedAgent.avatar,
            voice: updatedAgent.voice,
            metrics: updatedAgent.metrics,
            recordings: updatedAgent.recordings
          }
        }
      })
      setAgent(updatedAgent)
      toast.success('Agent updated successfully')
    } catch (_error) {
      toast.error('Failed to update agent')
    }
  }

  if (isLoading) {
    return <div>Loading...</div>
  }

  if (!agent) {
    return <div>Agent not found</div>
  }

  return (
    <>
      <Header>
        <div className="flex items-center gap-4">
          <Button
            variant="ghost"
            size="icon"
            onClick={() => navigate({ to: '/agents' })}
          >
            <IconArrowLeft className="h-5 w-5" />
          </Button>
          <h1 className="text-lg font-medium">{agent.name}</h1>
        </div>
        <div className="flex items-center gap-4">
          <ThemeSwitch />
          <ProfileDropdown />
        </div>
      </Header>
      <Main>
        <AgentDetails agent={agent} onUpdate={handleUpdate} />
      </Main>
    </>
  )
}
