import { Button } from "@/components/ui/button";
import { Card } from "@/components/ui/card";
import { IconPhoto, IconX } from "@tabler/icons-react";
import { useCallback, useState } from "react";
import { useDropzone } from "react-dropzone";

interface ImageUploadProps {
  onChange: (file: File | null) => void;
  value?: string;
  className?: string;
}

export function ImageUpload({ onChange, value, className }: ImageUploadProps) {
  const [preview, setPreview] = useState<string | null>(value || null);

  const onDrop = useCallback(
    (acceptedFiles: File[]) => {
      if (acceptedFiles.length > 0) {
        const file = acceptedFiles[0];
        onChange(file);
        const objectUrl = URL.createObjectURL(file);
        setPreview(objectUrl);
      }
    },
    [onChange]
  );

  const { getRootProps, getInputProps, isDragActive } = useDropzone({
    onDrop,
    accept: {
      "image/*": [".png", ".jpg", ".jpeg", ".gif"],
    },
    maxFiles: 1,
  });

  const handleRemove = () => {
    onChange(null);
    setPreview(null);
  };

  return (
    <Card
      className={`relative border-2 border-dashed rounded-lg ${
        isDragActive ? "border-primary" : "border-border"
      } ${className}`}
    >
      {preview ? (
        <div className="relative aspect-square w-full overflow-hidden rounded-lg">
          <img
            src={preview}
            alt="Preview"
            className="object-cover w-full h-full"
          />
          <Button
            type="button"
            variant="destructive"
            size="icon"
            className="absolute top-2 right-2"
            onClick={handleRemove}
          >
            <IconX className="h-4 w-4" />
          </Button>
        </div>
      ) : (
        <div
          {...getRootProps()}
          className="flex flex-col items-center justify-center p-6 cursor-pointer"
        >
          <input {...getInputProps()} />
          <IconPhoto className="h-10 w-10 text-muted-foreground mb-2" />
          <p className="text-sm text-muted-foreground text-center">
            {isDragActive
              ? "Drop the image here"
              : "Drag & drop an image here, or click to select"}
          </p>
        </div>
      )}
    </Card>
  );
}
