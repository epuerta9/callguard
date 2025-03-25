import React, { createContext, useContext, useState } from "react";

// Define a type for the context value
interface SettingsPropsContextType {
  formikHelpers: any;
  setLogoFile: (event: any) => void;
  setGalleryImgFiles: (event: any) => void;
  galleryImgUrls: string[];
  setGalleryImgUrls: (event: any) => void;
  themeColors: string[];
  setThemeColors: (event: any) => void;
}

const SettingsPropsContext = createContext<
  SettingsPropsContextType | undefined
>(undefined);

export const SettingsStoreProvider: React.FC<{
  children: React.ReactNode;
  formikHelpers: any;
  setLogoFile: (event: any) => void;
  setGalleryImgFiles: (event: any) => void;
  galleryImgUrls: string[];
  setGalleryImgUrls: (event: any) => void;
  themeColors: string[];
  setThemeColors: (event: any) => void;
}> = ({
  children,
  formikHelpers,
  setLogoFile,
  setGalleryImgFiles,
  galleryImgUrls,
  setGalleryImgUrls,
  themeColors,
  setThemeColors,
}) => {
  return (
    <SettingsPropsContext.Provider
      value={{
        formikHelpers,
        setLogoFile,
        setGalleryImgFiles,
        galleryImgUrls,
        setGalleryImgUrls,
        themeColors,
        setThemeColors,
      }}
    >
      {children}
    </SettingsPropsContext.Provider>
  );
};

export const useSettingStore = (): SettingsPropsContextType => {
  const context = useContext(SettingsPropsContext);
  if (!context) {
    throw new Error("useStore must be used within a StoreProvider");
  }
  return context;
};

interface SettingNavContextType {
  navSelected: boolean;
  setNavSelected: (event: any) => void;
}

const SettingNavContext = createContext<SettingNavContextType | undefined>(
  undefined
);

export const SettingNavStoreProvider: React.FC<{
  children: React.ReactNode;
}> = ({ children }) => {
  const [navSelected, setNavSelected] = useState(false);

  return (
    <SettingNavContext.Provider value={{ navSelected, setNavSelected }}>
      {children}
    </SettingNavContext.Provider>
  );
};

export const useSettingNavStore = (): SettingNavContextType => {
  const context = useContext(SettingNavContext);
  if (!context) {
    throw new Error("useStore must be used within a StoreProvider");
  }
  return context;
};

export { SettingNavContext, SettingsPropsContext };
