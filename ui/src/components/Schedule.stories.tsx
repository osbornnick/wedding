import type { Meta, StoryObj } from '@storybook/react-vite'
import Schedule from './Schedule'

const meta = {
  title: 'Wedding/Schedule',
  component: Schedule,
  parameters: {
    layout: 'fullscreen',
  },
  tags: ['autodocs'],
} satisfies Meta<typeof Schedule>

export default meta
type Story = StoryObj<typeof meta>

export const Default: Story = {}
