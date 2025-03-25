import { createLazyFileRoute } from '@tanstack/react-router'
import { BusinessBasicsForm } from '@/features/settings/business/basics-form'

export const Route = createLazyFileRoute(
  '/_authenticated/settings/business/basics'
)({
  component: BusinessBasicsForm,
})
