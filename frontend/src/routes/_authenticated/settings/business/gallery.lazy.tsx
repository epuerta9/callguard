import { createLazyFileRoute } from '@tanstack/react-router'
import { BusinessGalleryForm } from '@/features/settings/business/gallery-form'

export const Route = createLazyFileRoute(
  '/_authenticated/settings/business/gallery'
)({
  component: BusinessGalleryForm,
})
