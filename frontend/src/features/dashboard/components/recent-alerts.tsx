import { Avatar, AvatarFallback } from '@/components/ui/avatar'
import { AlertTriangle, ShieldAlert } from 'lucide-react'

export function RecentAlerts() {
  return (
    <div className='space-y-8'>
      {recentAlerts.map((alert) => (
        <div key={alert.id} className='flex items-center'>
          <Avatar className='h-9 w-9'>
            <AvatarFallback className={alert.type === 'spam' ? 'bg-yellow-500/10 text-yellow-500' : 'bg-red-500/10 text-red-500'}>
              {alert.type === 'spam' ? <AlertTriangle className='h-4 w-4' /> : <ShieldAlert className='h-4 w-4' />}
            </AvatarFallback>
          </Avatar>
          <div className='ml-4 space-y-1'>
            <p className='text-sm font-medium leading-none'>{alert.type === 'spam' ? 'Spam Call' : 'Phishing Attempt'}</p>
            <p className='text-sm text-muted-foreground'>
              Agent: {alert.agent} • {alert.time}
            </p>
          </div>
          <div className='ml-auto font-medium'>
            {alert.confidence}% confidence
          </div>
        </div>
      ))}
    </div>
  )
}

const recentAlerts = [
  {
    id: '1',
    type: 'phishing',
    agent: 'Sarah Miller',
    time: '2 mins ago',
    confidence: 98,
  },
  {
    id: '2',
    type: 'spam',
    agent: 'John Davis',
    time: '5 mins ago',
    confidence: 95,
  },
  {
    id: '3',
    type: 'phishing',
    agent: 'Emma Wilson',
    time: '12 mins ago',
    confidence: 92,
  },
  {
    id: '4',
    type: 'spam',
    agent: 'Michael Brown',
    time: '25 mins ago',
    confidence: 89,
  },
  {
    id: '5',
    type: 'phishing',
    agent: 'David Lee',
    time: '32 mins ago',
    confidence: 97,
  },
]
