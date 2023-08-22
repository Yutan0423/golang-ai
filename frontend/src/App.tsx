import { Suspense } from "react";
import "./App.css";
import { Body } from "./components/Body";
import { Button } from "./components/Button";
import { Layout } from "./components/Layout";
import { useFetchAi } from "./hooks/useFetchAi";

export default function App() {
  const { data, isLoading } = useFetchAi();
  console.log(data);

  const handleSubmit = () => {
    
  }

  return (
    <Suspense fallback={<div>Loading...</div>}>
      <Layout>
        <h1 className="text-3xl font-bold">設計力診断</h1>
        {isLoading ? <div>Loading...</div> : <Body data={data} />}
        <Button onClick={handleSubmit}>DBに登録する</Button>
      </Layout>
    </Suspense>
  );
}
