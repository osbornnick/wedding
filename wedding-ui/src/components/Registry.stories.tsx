import type { Meta, StoryObj } from '@storybook/react-vite'
import Registry, { defaultRegistries } from './Registry'

const meta = {
  title: 'Wedding/Registry',
  component: Registry,
  parameters: {
    layout: 'fullscreen',
  },
  tags: ['autodocs'],
} satisfies Meta<typeof Registry>

export default meta
type Story = StoryObj<typeof meta>

export const Default: Story = {
  args: {
    items: defaultRegistries,
  },
}

export const SingleItem: Story = {
  args: {
    items: [
      {
        name: 'Zola',
        url: 'https://www.zola.com',
        description: 'Our full gift registry and honeymoon fund.',
      },
    ],
  },
}
