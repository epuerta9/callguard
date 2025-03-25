import { useState } from 'react'
import { useNavigate } from '@tanstack/react-router'
import {
  IconPlus,
  IconSearch,
  IconUser,
  IconPhone,
  IconSettings,
  IconAdjustmentsHorizontal,
  IconSortAscendingLetters,
  IconSortDescendingLetters,
} from '@tabler/icons-react'
import { Button } from '@/components/ui/button'
import { Input } from '@/components/ui/input'
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogFooter,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from '@/components/ui/dialog'
import { Label } from '@/components/ui/label'
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from '@/components/ui/select'
import { Separator } from '@/components/ui/separator'
import { Header } from '@/components/layout/header'
import { Main } from '@/components/layout/main'
import { ProfileDropdown } from '@/components/profile-dropdown'
import { ThemeSwitch } from '@/components/theme-switch'
import { Avatar, AvatarFallback, AvatarImage } from '@/components/ui/avatar'
import { Badge } from '@/components/ui/badge'
import { Card, CardContent, CardHeader } from '@/components/ui/card'
import { agents, type Agent } from './data/agents'

const agentTypeText = new Map<string, string>([
  ['all', 'All Agents'],
  ['general', 'General'],
  ['custom', 'Custom'],
])

export default function Agents() {
  const navigate = useNavigate()
  const [sort, setSort] = useState('ascending')
  const [agentType, setAgentType] = useState('all')
  const [searchTerm, setSearchTerm] = useState('')
  const [isCreateOpen, setIsCreateOpen] = useState(false)
  const [newAgent, setNewAgent] = useState({
    name: '',
    phoneNumber: ''
  })

  const filteredAgents = [...agents]
    .sort((a: Agent, b: Agent) =>
      sort === 'ascending'
        ? a.name.localeCompare(b.name)
        : b.name.localeCompare(a.name)
    )
    .filter((agent: Agent) =>
      agentType === 'general'
        ? agent.type === 'General'
        : agentType === 'custom'
          ? agent.type === 'Custom'
          : true
    )
    .filter((agent: Agent) =>
      agent.name.toLowerCase().includes(searchTerm.toLowerCase()) ||
      agent.phoneNumber.includes(searchTerm)
    )

  const handleCreateAgent = () => {
    // TODO: Implement agent creation
    setIsCreateOpen(false)
    setNewAgent({ name: '', phoneNumber: '' })
  }

  return (
    <>
      <Header>
        <div className="flex items-center gap-4">
          <h1 className="text-2xl font-bold">Agents</h1>
          <Badge variant="secondary">Dashboard / Agents</Badge>
        </div>
        <div className='ml-auto flex items-center gap-4'>
          <ThemeSwitch />
          <ProfileDropdown />
        </div>
      </Header>

      <Main>
        <div className="flex items-center justify-between mb-6">
          <div>
            <h2 className="text-xl font-semibold">Your agents</h2>
            <p className="text-muted-foreground">Click one of your agent to see its setting or analytics</p>
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
              <form id="agent-form" onSubmit={(e) => {
                e.preventDefault();
                handleCreateAgent();
              }}>
                <div className="grid gap-4 py-4">
                  <div className="grid gap-2">
                    <Label htmlFor="name">Name</Label>
                    <Input
                      id="name"
                      value={newAgent.name}
                      onChange={(e) => setNewAgent(prev => ({ ...prev, name: e.target.value }))}
                      placeholder="Enter agent name"
                    />
                  </div>
                  <div className="grid gap-2">
                    <Label htmlFor="phone">Phone Number</Label>
                    <Input
                      id="phone"
                      value={newAgent.phoneNumber}
                      onChange={(e) => setNewAgent(prev => ({ ...prev, phoneNumber: e.target.value }))}
                      placeholder="(XXX) XXX-XXXX"
                    />
                  </div>
                </div>
                <DialogFooter>
                  <Button variant="outline" onClick={() => setIsCreateOpen(false)}>Cancel</Button>
                  <Button type='submit'>Save changes</Button>
                </DialogFooter>
              </form>
            </DialogContent>
          </Dialog>
        </div>

        <div className="mb-6">
          <div className="relative">
            <IconSearch className="absolute left-3 top-1/2 -translate-y-1/2 text-muted-foreground" size={18} />
            <Input
              placeholder="Search agents..."
              value={searchTerm}
              onChange={(e) => setSearchTerm(e.target.value)}
              className="pl-10 w-[300px]"
            />
          </div>
        </div>

        <div className="grid gap-4 md:grid-cols-2 lg:grid-cols-3">
          {filteredAgents.map((agent) => (
            <Card
              key={agent.id}
              className="cursor-pointer hover:shadow-md transition-shadow"
              onClick={() => navigate({ to: `/agents/${agent.id}` })}
            >
              <CardHeader className="pb-4">
                <div className="flex items-center gap-3">
                  <Avatar>
                    <AvatarImage src={agent.avatar} />
                    <AvatarFallback>{agent.name[0]}</AvatarFallback>
                  </Avatar>
                  <div>
                    <h3 className="font-semibold">{agent.name}</h3>
                    <p className="text-sm text-muted-foreground">{agent.phoneNumber}</p>
                  </div>
                </div>
              </CardHeader>
              <CardContent>
                <div className="flex gap-2">
                  <Badge variant="secondary" className="flex items-center gap-1">
                    <IconUser size={14} />
                    {agent.type}
                  </Badge>
                  <Badge variant="outline" className="flex items-center gap-1">
                    <IconPhone size={14} />
                    {agent.status}
                  </Badge>
                </div>
              </CardContent>
            </Card>
          ))}
        </div>
      </Main>
    </>
  )
}
