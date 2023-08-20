import { FC, ReactNode } from "react";

type Props = {
  children: ReactNode;
};

export const Layout: FC<Props> = ({ children }) => (
  <div className="flex flex-col items-center justify-center min-h-screen bg-gray-100">
    <main className="flex flex-col items-center justify-center flex-1 w-full px-20 text-center">
      {children}
    </main>
  </div>
);
