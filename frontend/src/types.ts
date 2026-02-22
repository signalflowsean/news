export interface Story {
  id: string;
  title: string;
  summary?: string;
  createdAt: string;
}

export interface InitialData {
  stories: Story[];
}

declare global {
  interface Window {
    __INITIAL_DATA__?: InitialData;
  }
}
