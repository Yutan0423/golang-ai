import { FC } from "react";
import { Question } from "../../domains/openai";

type Props = {
  data: Question[];
};

export const Body: FC<Props> = ({ data }) => (
  <div>
    {data.length > 0 &&
      data.map((q) => (
        <div key={q.question} className="my-6">
          <p className="text-2xl font-bold">{q.question}</p>
          <ul className="p-4">
            {q.choices &&
              q.choices.length > 0 &&
              q.choices.map((choice, idx) => <li key={idx}>
                  <span>
                  {choice}
                  </span>
                </li>)}
          </ul>
          <p>答え: {q.answer}</p>
        </div>
      ))}
  </div>
);
