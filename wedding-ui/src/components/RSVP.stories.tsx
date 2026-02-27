import type { Meta, StoryObj } from '@storybook/react-vite'
import RSVP from './RSVP'

const meta = {
  title: 'Wedding/RSVP',
  component: RSVP,
  parameters: {
    layout: 'fullscreen',
  },
  tags: ['autodocs'],
} satisfies Meta<typeof RSVP>

export default meta
type Story = StoryObj<typeof meta>

export const Form: Story = {
  args: {
    initialSubmitted: false,
  },
}

export const Confirmed: Story = {
  args: {
    initialSubmitted: true,
  },
}
