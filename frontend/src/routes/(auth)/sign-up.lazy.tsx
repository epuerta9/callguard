import { createLazyFileRoute } from '@tanstack/react-router'
import SignUp from '@/features/auth/sign-up/sign-up'

export const Route = createLazyFileRoute('/(auth)/sign-up')({
  component: SignUp,
})
