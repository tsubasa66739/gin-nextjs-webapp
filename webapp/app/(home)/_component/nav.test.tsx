import { render } from "@testing-library/react";
import Nav from "./nav";

jest.mock("next/navigation", () => ({
  useRouter: () => jest.fn(),
  usePathname: () => jest.fn().mockReturnValue("/"),
}));

describe("nav render", () => {
  test("Nav 正常", async () => {
    const { getByRole } = render(<Nav />);

    expect(getByRole("button", { name: "ホーム" })).toBeInTheDocument();
    expect(getByRole("button", { name: "カレンダー" })).toBeInTheDocument();
    expect(getByRole("button", { name: "データ" })).toBeInTheDocument();
    expect(getByRole("button", { name: "設定" })).toBeInTheDocument();
  });
});
