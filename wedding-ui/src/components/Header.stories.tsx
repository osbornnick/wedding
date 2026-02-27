import type { Meta, StoryObj, StoryFn } from '@storybook/react-vite'
import {
  RouterProvider,
  createRootRoute,
  createRouter,
  createMemoryHistory,
  Outlet,
} from '@tanstack/react-router'
import Header from './Header'

const rootRoute = createRootRoute({ component: Outlet })
const router = createRouter({
  routeTree: rootRoute,
  history: createMemoryHistory(),
})

const withRouter = (Story: StoryFn) => (
  <RouterProvider router={router}>
    <Story />
  </RouterProvider>
)

const meta = {
  title: 'Wedding/Header',
  component: Header,
  parameters: {
    layout: 'fullscreen',
  },
  decorators: [withRouter],
  tags: ['autodocs'],
} satisfies Meta<typeof Header>

export default meta
type Story = StoryObj<typeof meta>

export const Default: Story = {}
