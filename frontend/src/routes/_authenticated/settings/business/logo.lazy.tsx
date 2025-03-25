import { createLazyFileRoute } from '@tanstack/react-router'
import { BusinessLogoForm } from '@/features/settings/business/logo-form'

export const Route = createLazyFileRoute(
  '/_authenticated/settings/business/logo'
)({
  component: BusinessLogoForm,
})
