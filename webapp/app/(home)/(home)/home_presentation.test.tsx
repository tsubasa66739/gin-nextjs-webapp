import { render } from "@testing-library/react";
import HomePresentation from "./home_presentation";

describe("home render", () => {
  test("ホーム画面 正常", () => {
    const { getByRole } = render(
      <HomePresentation
        notes={[
          {
            id: 1,
            title: "fuga",
            body: "hoge",
            createdAt: "2024-12-16 13:08:03.239819+00",
            updatedAt: "2024-12-16 13:08:03.239819+00",
          },
        ]}
      />
    );

    expect(getByRole("button", { name: "Other control" })).toBeInTheDocument();
    expect(getByRole("button", { name: "New Note" })).toBeInTheDocument();

    expect(getByRole("columnheader", { name: "タイトル" })).toBeInTheDocument();
    expect(getByRole("columnheader", { name: "本文" })).toBeInTheDocument();
    expect(getByRole("columnheader", { name: "更新日時" })).toBeInTheDocument();
    expect(getByRole("columnheader", { name: "" })).toBeInTheDocument();

    expect(getByRole("rowheader", { name: "fuga" })).toBeInTheDocument();
    expect(getByRole("gridcell", { name: "hoge" })).toBeInTheDocument();
    expect(
      getByRole("gridcell", { name: "2024/12/16 22:08" })
    ).toBeInTheDocument();
  });

  test("ホーム画面 データなし", () => {
    const { getByRole } = render(<HomePresentation notes={[]} />);

    expect(getByRole("columnheader", { name: "タイトル" })).toBeInTheDocument();
    expect(getByRole("columnheader", { name: "本文" })).toBeInTheDocument();
    expect(getByRole("columnheader", { name: "更新日時" })).toBeInTheDocument();
    expect(getByRole("columnheader", { name: "" })).toBeInTheDocument();

    expect(getByRole("gridcell", { name: "No data." })).toBeInTheDocument();
  });
});
