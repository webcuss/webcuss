export const isCtrlEnter = (e: React.KeyboardEvent<HTMLElement>) => (e.ctrlKey || e.metaKey) && e.key == "Enter";
