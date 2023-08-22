import useSWR from "swr";
import { Question } from "../domains/openai";
import { fetchAi } from "../interfaces/openai";

type Response = {
  data: Question[];
  isLoading: boolean;
  isError: boolean;
};

export const useFetchAi = (): Response => {
  const { data, error } = useSWR("/api/ai", fetchAi);

  return {
    data: data ?? [],
    isLoading: !error && !data,
    isError: !!error,
  };
};
