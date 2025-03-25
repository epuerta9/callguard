import { useParams } from '@tanstack/react-router'
import { agents } from '@/features/agents/data/agents'
import { AgentDetails } from '@/features/agents/components/agent-details'
import { Main } from '@/components/layout/main'
import { Header } from '@/components/layout/header'
import { ProfileDropdown } from '@/components/profile-dropdown'
import { ThemeSwitch } from '@/components/theme-switch'
import { Button } from '@/components/ui/button'
import { IconArrowLeft } from '@tabler/icons-react'
import { useNavigate } from '@tanstack/react-router'

export default function AgentDetailsPage() {
  const { id } = useParams({ from: '/_authenticated/agents/[id]' })
  const navigate = useNavigate()
  const agent = agents.find((a) => a.id === id)

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
            onClick={() => navigate({ to: '/' })}
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
        <AgentDetails agent={agent} onUpdate={() => {}} />
      </Main>
    </>
  )
}
