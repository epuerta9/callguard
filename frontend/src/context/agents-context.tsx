// import React, { createContext, useContext, ReactNode } from "react";
// import { useAgents } from "../hooks/useAgents";
// import { agentData } from "../components/Agents/types";

// interface AgentsContextProps {
//   creditsSum: number;
//   agentsDataArr: agentData[];
//   refreshAgents: () => void;
// }

// const initialState: AgentsContextProps = {
//   creditsSum: 0,
//   agentsDataArr: [],
//   refreshAgents: () => {},
// };

// const AgentsContext = createContext<AgentsContextProps>(initialState);

// export const AgentsProvider: React.FC<{ children: ReactNode }> = ({
//   children,
// }) => {
//   const { agentsDataArr, creditsSum, refreshAgents } = useAgents();

//   return (
//     <AgentsContext.Provider
//       value={{ agentsDataArr, creditsSum, refreshAgents }}
//     >
//       {children}
//     </AgentsContext.Provider>
//   );
// };

// export const useAgentsStore = () => {
//   const context = useContext(AgentsContext);
//   if (!context) {
//     throw new Error("useAgentsStore must be used within an AgentsProvider");
//   }
//   return context;
// };

// export default AgentsContext;
