import { createLazyFileRoute } from '@tanstack/react-router'
import { BusinessHoursForm } from '@/features/settings/business/hours-form'

export const Route = createLazyFileRoute(
  '/_authenticated/settings/business/hours'
)({
  component: BusinessHoursForm,
})
