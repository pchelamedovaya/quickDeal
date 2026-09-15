export interface CreateAdRequest {
  title: string;
  description: string;
  price: number;
}

export interface AdResponse {
  id: string;
  title: string;
  description: string;
  price: number;
  author: string;
  created_at: string;
}
