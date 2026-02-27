import type { Meta, StoryObj } from '@storybook/react-vite'
import Travel from './Travel'

const meta = {
  title: 'Wedding/Travel',
  component: Travel,
  parameters: {
    layout: 'fullscreen',
  },
  tags: ['autodocs'],
} satisfies Meta<typeof Travel>

export default meta
type Story = StoryObj<typeof meta>

export const Default: Story = {}
