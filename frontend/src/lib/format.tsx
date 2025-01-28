/* eslint-disable no-irregular-whitespace */
import clsx from "clsx";
import { ArrowDown, ArrowUp, Minus } from "lucide-react";

export function formatToTwoDecimalsPrice(value: number): string {
  const floatValue = value / 1_000_000;
  // eslint-disable-next-line no-irregular-whitespace
  return `${floatValue.toFixed(2)} €`;
}

export function formatToPercentage(value: number): JSX.Element {
  return (
    <div
      className={clsx(
        "flex tabular-nums",
        value > 0
          ? "text-green-600"
          : value < 0
          ? "text-red-600"
          : "text-blue-600"
      )}
    >
      {value > 0 ? (
        <ArrowUp className="text-sm" />
      ) : value < 0 ? (
        <ArrowDown className="text-sm" />
      ) : (
        <Minus className="text-sm" />
      )}
       {value.toFixed(2).replace(/-/g, "")}%
    </div>
  );
}
