import { createFileRoute } from '@tanstack/react-router'

export const Route = createFileRoute('/_authenticated/agents/[id]')({
  beforeLoad: ({ params }: { params: { id: string } }) => ({ id: params.id }),
  component: () =>
    import('./[id].lazy').then((mod) => ({
      default: mod.default,
    })),
})
