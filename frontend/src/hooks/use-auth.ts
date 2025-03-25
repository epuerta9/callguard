import { useNavigate } from '@tanstack/react-router'
import { toast } from '@/hooks/use-toast'
import { useEffect } from 'react'
import { IconCircleCheck } from '@tabler/icons-react'
import React from 'react'
import useApi from '@/api/api'

interface User {
  access_token: string
  user_id: string
  account_id: string
  first_name?: string
  last_name?: string
  email?: string
  avatar_url?: string
}

export function useAuth() {
  const navigate = useNavigate()
  const [isLoadingProfile, setIsLoadingProfile] = React.useState(false)
  const [user, setUser] = React.useState<User | null>(() => {
    const userStr = localStorage.getItem('user')
    if (!userStr) return null
    try {
      return JSON.parse(userStr)
    } catch {
      return null
    }
  })

  // Create a stable API instance
  const api = React.useMemo(() => useApi(), []);

  const hasFetchedRef = React.useRef(false)

  useEffect(() => {
    if (hasFetchedRef.current || !user?.user_id || user?.first_name) return
    hasFetchedRef.current = true

    const fetchUserData = async () => {
      try {
        setIsLoadingProfile(true)
        const response = await api.get(`/users/${user.user_id}`)
        if (response?.data) {
          const updatedUser = { ...user, ...response.data }
          localStorage.setItem('user', JSON.stringify(updatedUser))
          setUser(updatedUser)
        }
      } catch (err) {
        console.error('Failed to fetch user data:', err)
      } finally {
        setIsLoadingProfile(false)
      }
    }

    fetchUserData()
  }, [user?.user_id, api]) // cleaner dependency array


  const logout = React.useCallback(() => {
    // Clear user data
    localStorage.removeItem('user')
    setUser(null)
    toast({
      title: 'Logged out',
      description: 'You have been logged out successfully',
      variant: 'default'
    })
    window.location.href = '/sign-in'
  }, [])

  const validateToken = () => {
    if (!user?.token) {
      logout()
      return false
    }

    // Check if token is expired by decoding JWT
    try {
      const payload = JSON.parse(atob(user?.token.split('.')[1]))
      if (payload.exp * 1000 < Date.now()) {
        logout()
        return false
      }
    } catch {
      logout()
      return false
    }

    return true
  }

  // Validate token on mount and setup interval to check periodically
  useEffect(() => {
    // Only set up validation if we have a user
    if (!user) return

    validateToken()
    const interval = setInterval(validateToken, 60000) // Check every minute
    return () => clearInterval(interval)
  }, [])

  return {
    user,
    setUser,
    logout,
    validateToken
  }
}
