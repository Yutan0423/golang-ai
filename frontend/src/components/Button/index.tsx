import { FC, ReactNode } from "react";

type Props = {
    onClick: () => void;
    children: ReactNode;
}

export const Button: FC<Props> = ({ onClick, children }) => (
    <button onClick={onClick} className="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded">
        {children}
    </button>
)