export const insertQuestions = async (): Promise<void> => {
    const response = await fetch("/api/ai").then((res) => res.json());
    return response;
  };