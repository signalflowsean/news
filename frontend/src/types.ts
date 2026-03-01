export interface Story {
  id: string;
  title: string;
  summary?: string;
  createdAt: string;
}

export interface InitialData {
  stories: Story[];
}
