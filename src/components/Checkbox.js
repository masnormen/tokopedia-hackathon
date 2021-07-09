export const Checkbox = ({ label }) => {
  return (
    <label className="flex items-center cursor-pointer my-1.5 w-full">
      <input
        type="checkbox"
        className="form-checkbox h-5 w-5 border-2 border-gray-500 rounded-md text-xgreen"
      />
      <span className="ml-2 rounded-md text-xs text-gray-500">{label}</span>
    </label>
  );
};
