import { Link, useNavigate } from '@tanstack/react-router'
import {
  DarkThemeToggle,
  Button,
  Navbar,
  NavbarBrand,
  NavbarCollapse,
  NavbarToggle,
  NavbarLink,
} from 'flowbite-react'

const navLinks = [
  { to: '/', label: 'Home', exact: true },
  { to: '/schedule', label: 'Schedule' },
  { to: '/travel', label: 'Travel' },
  { to: '/rsvp', label: 'RSVP' },
  { to: '/registry', label: 'Registry' },
] as const

export default function Header() {
  const navigate = useNavigate()

  async function handleLogout() {
    navigate({ to: '/logout' })
  }

  return (
    <header className="sticky top-0 z-50 w-full border-b border-gray-200 dark:border-gray-700">
      <Navbar fluid className="max-w-6xl mx-auto">
        <NavbarBrand as={Link} href="/">
          <span className="font-serif text-xl text-gray-700 hover:bg-gray-50 md:border-0 md:hover:bg-transparent md:hover:text-primary-700 dark:border-gray-700 dark:text-gray-400 dark:hover:bg-gray-700 dark:hover:text-white md:dark:hover:bg-transparent md:dark:hover:text-white">
            Nicky &amp; Sarah
          </span>
        </NavbarBrand>
        <NavbarToggle />
        <NavbarCollapse>
          {navLinks.map(({ to, label, exact }) => (
            <NavbarLink key={to} href={to}>
              {label}
            </NavbarLink>
          ))}
        </NavbarCollapse>
        <div className="flex">
          {import.meta.env.DEV && (
            <Button onClick={handleLogout}>Dev: Logout</Button>
          )}
          <DarkThemeToggle />
        </div>
      </Navbar>
    </header>
  )
}
