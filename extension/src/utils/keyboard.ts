export const isCtrlEnter = (e: React.KeyboardEvent<HTMLTextAreaElement | HTMLInputElement>) => (e.ctrlKey || e.metaKey) && e.key == "Enter";
