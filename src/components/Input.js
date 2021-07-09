export const Input = ({ label, ...rest }) => (
  <div>
    <label className="block font-semibold text-rose-900 text-md mb-2">
      {label}
    </label>
    <input
      className="shadow appearance-none border border-gray-200 rounded-md w-full px-4 py-3 focus:outline-none focus:shadow-lg"
      {...rest}
    />
  </div>
);
