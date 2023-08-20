import { Question } from "../domains/ai";

export const fetchAi = async (): Promise<Question[]> => {
  const response = await fetch("/api/ai").then((res) => res.json());
  return response;
};
